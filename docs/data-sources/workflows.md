---
subcategory: "Workflow"
layout: "opswat"
page_title: "OPSWAT: opswat_workflows"
sidebar_current: "docs-opswat-datasource-file-sync"
description: |-
  OPSWAT Workflows.
---

## Example Usage

```terraform
data "opswat_workflows" "current" {}

output "opswat_workflows" {
  value = data.opswat_workflows.current
}
```

## Schema
Read-only:  
- `allow_cert` - (String) Generate batch signature with certificate - Use certificate to generate batch signature flag.
- `allow_cert_cert` - (String) Certificate used for barch signing.
- `allow_cert_cert_validity` - (Int) Certificate validity (hours).
- `allow_local_files` - (Bool)  Process files from servers - Allow scan on server flag.
- `allow_local_files_local_paths` - (List of strings) Server scan local paths.
- `allow_local_files_white_list` - (String) Process files from servers flag (false = ALLOW ALL EXCEPT, true = DENY ALL EXCEPT).
- `description` - (String) Workflow description.
- `id` - (Int) Workflow ID.
- `include_webhook_signature` - (String) Webhook - Include webhook signature flag.
- `include_webhook_signature_certificate_id` - (Int) Webhook signature certificate id.
- `last_modified` - (Int) Last modified timestamp (unix epoch).
- `mutable` - (Bool) Mutable flag.
- `name` - (String) Workflow name.
- `option_values` - (Object)
  - `archive_handling_max_number_files` - (Int) Archive - Process - Max number of files extracted.
  - `archive_handling_max_recursion_level` - (Int) Archive - Process - Max recursion level.
  - `archive_handling_max_size_files` - (Int) Archive - Process - Max total size of extracted files.
  - `archive_handling_timeout` - (Int) Archive - Timeout - Archive analysis timeout.
  - `filetype_analysis_timeout` - (Int) File Type - Timeout - File type analysis timeout.
  - `process_info_global_timeout` - (Bool) General - Enable global timeout flag.
  - `process_info_global_timeout_value` - (Int) General - Global timeout (in seconds).
  - `process_info_max_download_size` - (Int) General - URL file download (in MB).
  - `process_info_max_file_size` - (Int) General - File scan (in MB).
  - `process_info_quarantine` - (Bool) General - Process - Quarantine blocked files flag.
  - `process_info_skip_hash` - (Bool) General - Process - Skip hash calculation flag.
  - `process_info_skip_processing_fast_symlink` - (Bool) General - Process - Skip processing fast symlink in archive flag.
  - `process_info_workflow_priority` - (Int) General - Quality of Service -  Priority [5 - Very high, 4 - High, 3 - Medium, 2 - Low, 1 - Very low].
- `scan_filescan_check_av_engine` - (Bool) Metascan - Thresholds - Require a min number of active AV engines for the whole file processing flag.
- `scan_filescan_download_timeout` - (Int) Metascan - Timeouts - File download (in minutes).
- `scan_filescan_global_scan_timeout` - (Int) Metascan - Timeouts - Global scan (in minutes).
- `scan_filescan_per_engine_scan_timeout` - (Int) Metascan - Timeouts - Per AV engine scan (in minutes).
- `vul_filescan_timeout_vulnerability_scanning` - (Int) Vulnerability - Timeouts - File-Based Vulnerability Assessment timeout (in minutes).
- `result_allowed` - (List of ints) Visibility of Processing result - Visibility of scan result.
  - `role` - (Int) Role ID.
  - `visibility` - (Int) Visibility scope ID [3 - FULL DETAILS, 2 - PER ENGINE RESULT, 1 - OVERALL RESULT ONLY]
- `scan_allowed` - (List of ints) Restrictions - Restrict access to following roles.
- `user_agents` - (List of strings) Restrictions - Limit to specified user agents.
- `workflow_id` - (Int) Workflow id.
- `zone_id` - (Int) Workflow network access zone id.