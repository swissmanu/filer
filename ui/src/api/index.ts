import { anArrayContaining, fromMap } from "spicery";
import { anInboxItem, InboxItem } from "../types/inboxItem";
import type { Rule } from "../types/rule";
import { aRule } from "../types/rule";

export class API {
  constructor(private readonly baseUrl: string = "") {}

  async getRules(): Promise<ReadonlyArray<Rule>> {
    const response = await fetch(`${this.baseUrl}/rules`);

    if (response.ok) {
      const json = await response.json();
      return fromMap(json, "rules", anArrayContaining(aRule));
    }

    throw new Error(
      `Could not process response with status code ${response.status} ${response.statusText}`
    );
  }

  async getInbox(): Promise<ReadonlyArray<InboxItem>> {
    const response = await fetch(`${this.baseUrl}/inbox`);

    if (response.ok) {
      const json = await response.json();
      return anArrayContaining(anInboxItem)(json);
    }

    throw new Error(
      `Could not process response with status code ${response.status} ${response.statusText}`
    );
  }
}
