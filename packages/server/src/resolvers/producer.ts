import { Arg, Int, Mutation, Query, Resolver } from "type-graphql";
import { getConnection } from "typeorm";
import { Producer } from "../entities/Producer";
import { Product } from "../entities/Product";
import { ProducerInput, UpdateProducerInput } from "./input/producer";

@Resolver()
export class ProducerResolver {
  @Mutation(() => Producer)
  createProducer(
    @Arg("input", () => ProducerInput) input: ProducerInput
  ): Promise<Producer> {
    return Producer.create(input).save();
  }

  @Query(() => Producer)
  producer(
    @Arg("uuid", () => String) uuid: string
  ): Promise<Producer | undefined> {
    return Producer.findOne({ uuid }, { relations: ["products"] });
  }

  @Query(() => [Producer])
  producers(): Promise<Producer[]> {
    return Producer.find({ relations: ["products"] });
  }

  @Mutation(() => Producer)
  async editProducer(
    @Arg("input", () => UpdateProducerInput) input: UpdateProducerInput,
    @Arg("uuid", () => String) uuid: number
  ): Promise<Producer | undefined> {
    const res = await getConnection()
      .createQueryBuilder()
      .update(Producer)
      .set({ name: input.name, canDeliver: input.canDeliver })
      .where("uuid = :uuid", { uuid })
      .returning("*")
      .execute();
    return res.raw[0];
  }

  @Mutation(() => Boolean)
  async deleteProducer(@Arg("id", () => Int) id: number): Promise<boolean> {
    await getConnection()
      .createQueryBuilder()
      .delete()
      .from(Product)
      .where('"ownerId" = :id', { id })
      .execute();

    await getConnection()
      .createQueryBuilder()
      .delete()
      .from(Producer)
      .where("id = :id", { id })
      .execute();

    return true;
  }
}
