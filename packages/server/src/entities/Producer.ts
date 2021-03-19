import { Field, Int, ObjectType } from "type-graphql";
import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from "typeorm";
import { Product } from "./Product";

@ObjectType()
@Entity("producers")
export class Producer extends BaseEntity {
  @Field(() => Int)
  @PrimaryGeneratedColumn()
  id!: number;

  @Field()
  @PrimaryGeneratedColumn("uuid")
  uuid!: string;

  @Field()
  @Column("text", { unique: true })
  name!: string;

  @Column("text")
  password!: string;

  @Field()
  @Column()
  canDeliver!: boolean;

  @Field(() => String)
  @CreateDateColumn()
  createdAt!: Date;

  // @Field()
  @OneToMany(() => Product, (product) => product.owner)
  products: Product[];
}
