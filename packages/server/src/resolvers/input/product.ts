import { Field, InputType, Int } from "type-graphql";

@InputType()
export class ProductInput {
  @Field()
  name!: string;

  @Field(() => Int)
  ownerId: number;

  @Field()
  ownerUuid: string;
}
