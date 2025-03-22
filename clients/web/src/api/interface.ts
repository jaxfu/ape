import * as Objects from "./impl/objects";
import * as Config from "./impl/config";

export const createObject = Objects.create;
export const getObjects = Objects.getAll;

export const pingServer = Config.pingServer;
