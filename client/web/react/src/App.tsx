import { useQuery } from "@tanstack/react-query";
import { GetComponents } from "./comps/Component";
import ApeComponent from "./comps/Component";

const App: React.FC = () => {
  const { isSuccess, data } = useQuery({
    queryKey: ["components"],
    queryFn: () => GetComponents(),
  });

  if (isSuccess && data !== undefined) {
    console.log(data);

    if (data.length === 0) {
      return <div>no components</div>;
    } else {
      return (
        <>
          {data.map((comp) => (
            <ApeComponent key={comp.component_id} component={comp} />
          ))}
        </>
      );
    }
  }

  return <h1>Ape</h1>;
};

export default App;
