import { App } from "cdktf";
import { CoreStack } from "./stacks/core";

const app = new App();
new CoreStack(app);
app.synth();
