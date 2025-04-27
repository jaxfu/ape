import { TComponentMetadata } from "../types";

const BASE_URL = "http://localhost:5000";

interface IProps {
  component: TComponentMetadata;
}

const ApeComponent: React.FC<IProps> = (props) => {
  return (
    <>
      <h2>{props.component.name}</h2>
      <div>{props.component.component_id}</div>
    </>
  );
};

export default ApeComponent;

export async function GetComponents(): Promise<TComponentMetadata[]> {
  const req = await fetch(`${BASE_URL}/api/components`);
  const res: TComponentMetadata[] = await req.json();

  return res;
}
