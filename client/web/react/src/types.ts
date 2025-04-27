export type TComponentMetadata = {
  name: string;
  component_id: string;
  component_type: EComponentType;
};

export type TProp = TComponentMetadata & {
  prop_metadata: TPropMetadata;
};

export type TPropMetadata = {
  prop_type: EPropType;
  is_array: boolean;
};

export type TPropConstraints = {};

export enum EComponentType {
  PROP = "PROP",
  OBJECT = "OBJECT",
  ROUTE = "ROUTE",
  MESSAGE_BODY = "MESSAGE_BODY",
  REQUEST = "REQUEST",
  RESPONSE = "RESPONSE",
}

export enum EPropType {
  INT = "INT",
  UINT = "UINT",
  FLOAT = "FLOAT",
  TEXT = "TEXT",
  BLOB = "BLOB",
  REF = "REF",
}
