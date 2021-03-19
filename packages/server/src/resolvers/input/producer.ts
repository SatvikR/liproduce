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

@InputType()
export class UpdateProducerInput {
  @Field({ nullable: true })
  name: string;

  @Field({ nullable: true })
  canDeliver: boolean;
}
