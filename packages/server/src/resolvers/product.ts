import { Arg, Mutation, Query, Resolver } from "type-graphql";
import { Product } from "../entities/Product";
import { ProductInput } from "./input/product";

@Resolver()
export class ProductResolver {
  @Mutation(() => Product)
  createProduct(@Arg("input", () => ProductInput) input: ProductInput) {
    return Product.create(input).save();
  }

  @Query(() => [Product])
  products() {
    return Product.find({ relations: ["owner"] });
  }
}
