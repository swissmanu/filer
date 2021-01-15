import { anArrayContaining, fromMap } from "spicery";
import { anInboxItem, InboxItem } from "../types/inboxItem";
import type { Rule } from "../types/rule";
import { aRule } from "../types/rule";

export class API {
  constructor(private readonly baseUrl: string = ".") {}

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

  async applyRuleToInboxItem(
    { name: ruleName }: Rule,
    { name: inboxItemName }: InboxItem,
    renameInboxItem?: string
  ): Promise<void> {
    const response = await fetch(
      `${this.baseUrl}/inbox/${inboxItemName}/apply`,
      {
        method: "post",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ruleName,
          ...(renameInboxItem ? { newInboxItemName: renameInboxItem } : {}),
        }),
      }
    );

    if (!response.ok) {
      throw new Error(
        `Could not apply rule! ${response.status} ${response.statusText}`
      );
    }
  }

  getUrlForItem({ name }: InboxItem): string {
    return `${this.baseUrl}/inbox/${name}`;
  }

  async deleteInboxItem({ name }: InboxItem): Promise<void> {
    const response = await fetch(`${this.baseUrl}/inbox/${name}`, {
      method: "delete",
    });

    if (!response.ok) {
      throw new Error(
        `Could not delete file. Response with status code ${response.status} ${response.statusText}`
      );
    }
  }
}
