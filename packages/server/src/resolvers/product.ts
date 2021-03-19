import { Arg, Mutation, Resolver } from "type-graphql";
import { Product } from "../entities/Product";
import { ProductInput } from "./input/product";

@Resolver()
export class ProductResolver {
  @Mutation(() => Product)
  createProduct(@Arg("input", () => ProductInput) input: ProductInput) {
    return Product.create(input).save();
  }
}
