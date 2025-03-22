import { BASE_URL } from "./config";
import { ApeProp } from "./props";
import { ComponentMetadata, PropType } from "./types";

type ApeObject = ComponentMetadata & {
  props?: { [key: string]: ApeProp };
};

export async function getAll(): Promise<Response> {
  return await fetch(`${BASE_URL}/object`, {
    method: "GET",
  });
}

const TEST_OBJ: ApeObject = {
  name: "Todo",
  props: {
    username: {
      type: PropType.TEXT,
    },
  },
};
export async function create(): Promise<Response> {
  return await fetch(`${BASE_URL}/object`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(TEST_OBJ),
  });
}
