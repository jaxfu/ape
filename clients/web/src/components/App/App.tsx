import { useEffect, useState } from "react";
import * as Api from "../../api/interface";

const App: React.FC = () => {
  const [healthy, setHealthy] = useState<boolean | null>(null);

  async function submit(): Promise<void> {
    const res = await Api.createObject();

    console.log(res.status === 201 ? "Object added" : "Error adding Object");
  }

  async function get(): Promise<void> {
    const res = await Api.getObjects();

    console.log(res.status);
    res.status === 200 && console.log(await res.json());
  }

  useEffect(() => {
    const fetchData = async () => {
      const result = await Api.pingServer();
      setHealthy(result);
    };
    fetchData();
  }, []);

  useEffect(() => {
    if (healthy !== null) {
      console.log(
        healthy
          ? "\u2713 Connected to server"
          : "\u274C Error connecting to server",
      );
    }
  }, [healthy]);

  return (
    <div>
      <button onClick={submit}>
        <h2>Send</h2>
      </button>

      <button onClick={get}>
        <h2>Get</h2>
      </button>
    </div>
  );
};

export default App;
