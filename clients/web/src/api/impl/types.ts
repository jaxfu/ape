export enum PropType {
  INT = "int",
  UINT = "uint",
  BOOL = "bool",
  TEXT = "text",
  FLOAT = "float",
  BLOB = "blob",
  REF = "ref",
}

export type ComponentMetadata = {
  name?: string;
  component_id?: string;
  display_id?: string;
  category?: string;
  description?: string;
};
