#!/usr/bin/env node

import cli from "commander";
import create_issue from "commands/issue/create_issue";
import delete_issue from "./commands/issue/delete_issue";
import update_issue from "./commands/issue/update_issue";
import find_issue from "./commands/issue/find_issue";

cli
  .command("create_issue")
  .argument("[id]", "ID of issue")
  .argument("[title]", "Title of issue")
  .argument("[description]", "Description of issue")
  .option("-p, --pretty", "Pretty-print output from the API.")
  .description(
    "Create issue"
  )
  .action(create_issue);

cli
  .command("delete_issue")
  .argument("[_id]", "ID of issue")
  .option("-p, --pretty", "Pretty-print output from the API.")
  .description(
    "Delete issue"
  )
  .action(delete_issue);

cli
  .command("update_issue")
  .argument("[id]", "ID of issue")
  .argument("[title]", "Title of issue")
  .argument("[description]", "Description of issue")
  .option("-p, --pretty", "Pretty-print output from the API.")
  .description(
    "Update issue"
  )
  .action(update_issue);

cli
  .command("find_issue")
  .argument("[_id]", "ID of issue")
  .option("-p, --pretty", "Pretty-print output from the API.")
  .description(
    "Find issue"
  )
  .action(find_issue);