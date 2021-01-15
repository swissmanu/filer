import {
  anArrayContaining,
  aString,
  fromMap,
  ParseFn,
  ParserError,
} from "spicery";

export interface Rule {
  name: string;
  description: string;
  actions: ReadonlyArray<Action>;
}

export const aRule: ParseFn<Rule> = (x) => ({
  name: fromMap(x, "name", aString),
  description: fromMap(x, "description", aString),
  actions: fromMap(x, "actions", anArrayContaining(anAction)),
});

export type Action = MoveAction;

export interface MoveAction {
  type: "move";
  target: string;
}

const anAction: ParseFn<Action> = (x) => {
  const type = fromMap(x, "type", aString);
  if (type !== "move") {
    throw new ParserError("ActionType", type);
  }

  return {
    type: "move",
    target: fromMap(x, "target", aString),
  };
};
