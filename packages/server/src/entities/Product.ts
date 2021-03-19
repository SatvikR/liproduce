import { Field, Int, ObjectType } from "type-graphql";
import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  ManyToOne,
  PrimaryGeneratedColumn,
} from "typeorm";
import { Producer } from "./Producer";

@ObjectType()
@Entity("products")
export class Product extends BaseEntity {
  @Field(() => Int)
  @PrimaryGeneratedColumn()
  id!: number;

  @Field()
  @PrimaryGeneratedColumn("uuid")
  uuid!: string;

  @Field()
  @Column("text")
  name!: string;

  @Field()
  @Column()
  ownerId!: number;

  @Field()
  @Column("uuid")
  ownerUuid!: string;

  @Field(() => Producer)
  @ManyToOne(() => Producer, (producer) => producer.products)
  owner!: Producer;

  @Field(() => String)
  @CreateDateColumn()
  createdAt!: Date;
}
