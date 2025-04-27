import { TProp } from "../types";

interface IProps {
  prop: TProp;
}

const ApeProp: React.FC<IProps> = (props) => {
  return (
    <>
      <h2>{props.prop.name}</h2>
    </>
  );
};
