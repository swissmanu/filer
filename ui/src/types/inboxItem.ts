import { aString, ParseFn } from "spicery";

export interface InboxItem {
  name: string;
}

export const anInboxItem: ParseFn<InboxItem> = (x) => {
  const name = aString(x);
  return { name };
};
