#!/usr/bin/env bash
# Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

defs_dir="${1:-generator/defs}"
docs_dir="${2:-docs/resources}"

if [[ ! -d "$defs_dir" ]]; then
    echo "Definitions directory not found: $defs_dir" >&2
    exit 1
fi

if [[ ! -d "$docs_dir" ]]; then
    echo "Resource documentation directory not found: $docs_dir" >&2
    exit 1
fi

updated=0
skipped=0

for yaml_file in "$defs_dir"/*.yaml; do
    [[ -f "$yaml_file" ]] || continue

    resource_name="$(awk '
        /^[[:space:]]{2}name:[[:space:]]*/ {
            value = $0
            sub(/^[[:space:]]*name:[[:space:]]*/, "", value)
            gsub(/^["'\''"]|["'\''"]$/, "", value)
            print value
            exit
        }
    ' "$yaml_file")"

    doc_category="$(awk '
        /^[[:space:]]{2}doc_category:[[:space:]]*/ {
            value = $0
            sub(/^[[:space:]]*doc_category:[[:space:]]*/, "", value)
            gsub(/^["'\''"]|["'\''"]$/, "", value)
            print value
            exit
        }
    ' "$yaml_file")"

    if [[ -z "$resource_name" || -z "$doc_category" ]]; then
        ((skipped += 1))
        continue
    fi

    doc_file="$docs_dir/${resource_name}.md"
    if [[ ! -f "$doc_file" ]]; then
        echo "Skipping $yaml_file: documentation file not found: $doc_file" >&2
        ((skipped += 1))
        continue
    fi

    tmp_file="$(mktemp)"
    awk -v subcategory="$doc_category" '
        BEGIN {
            frontmatter_delimiters = 0
            changed = 0
        }
        /^---[[:space:]]*$/ {
            frontmatter_delimiters += 1
        }
        frontmatter_delimiters == 1 && /^subcategory:[[:space:]]*/ {
            print "subcategory: \"" subcategory "\""
            changed = 1
            next
        }
        {
            print
        }
        END {
            if (changed == 0) {
                exit 2
            }
        }
    ' "$doc_file" > "$tmp_file" || {
        status=$?
        rm -f "$tmp_file"
        if [[ $status -eq 2 ]]; then
            echo "Skipping $doc_file: subcategory frontmatter field not found" >&2
            ((skipped += 1))
            continue
        fi
        exit "$status"
    }

    if ! cmp -s "$doc_file" "$tmp_file"; then
        mv "$tmp_file" "$doc_file"
        echo "Updated $doc_file -> subcategory: \"$doc_category\""
        ((updated += 1))
    else
        rm -f "$tmp_file"
        ((skipped += 1))
    fi
done

echo "Done. Updated: $updated, skipped: $skipped"
