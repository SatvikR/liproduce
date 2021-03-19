import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  ManyToOne,
  PrimaryGeneratedColumn,
} from "typeorm";
import { Producer } from "./Producer";

@Entity("products")
export class Product extends BaseEntity {
  @PrimaryGeneratedColumn()
  id!: number;

  @PrimaryGeneratedColumn("uuid")
  uuid!: string;

  @Column("text")
  name!: string;

  @Column()
  ownerId!: number;

  @Column("uuid")
  ownerUuid!: string;

  @ManyToOne(() => Producer, (producer) => producer.products)
  owner!: Producer;

  @CreateDateColumn()
  createdAt!: Date;
}
