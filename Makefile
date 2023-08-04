#!/usr/bin/env make

test-fixtures: tests/regions.json tests/expected.env

.PHONY: tests/regions.json # Forces rebuild
tests/regions.json:
	aws ssm get-parameters-by-path \
		--path /aws/service/global-infrastructure/regions \
		--no-cli-pager \
		--query 'sort_by(Parameters, &Name)[].{key:Value,value:Value}' | \
		jq -r '.[] | {key: .key | sub("-"; "_"; "g"), value: .value}' | \
		jq -s --indent 4 'from_entries' \
	> tests/regions.json

.PHONY: tests/expected.env # Forces rebuild
tests/expected.env:
	aws ssm get-parameters-by-path \
		--path /aws/service/global-infrastructure/regions \
		--no-cli-pager \
		--query 'sort_by(Parameters, &Name)[].{key:Value,value:Value}' | \
		jq -r '.[] | "\(.key | sub("-"; "_"; "g") | ascii_upcase)=\(.value)"' \
	> tests/expected.env
