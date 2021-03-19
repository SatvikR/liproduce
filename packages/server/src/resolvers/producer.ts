import { Arg, Mutation, Query, Resolver } from "type-graphql";
import { Producer } from "../entities/Producer";
import { ProducerInput } from "./input/producer";

@Resolver()
export class ProducerResolver {
  @Mutation(() => Producer)
  createProducer(
    @Arg("input", () => ProducerInput) input: ProducerInput
  ): Promise<Producer> {
    return Producer.create(input).save();
  }

  @Query(() => Producer)
  getOne(
    @Arg("uuid", () => String) uuid: string
  ): Promise<Producer | undefined> {
    return Producer.findOne({ uuid });
  }

  @Query(() => [Producer])
  getAll(): Promise<Producer[]> {
    return Producer.find();
  }
}
