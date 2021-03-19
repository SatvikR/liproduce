import { Field, InputType } from "type-graphql";

@InputType()
export class ProducerInput {
  @Field()
  name!: string;

  @Field()
  password!: string;

  @Field()
  canDeliver!: boolean;
}
