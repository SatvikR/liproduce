import {
  BaseEntity,
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from "typeorm";
import { Product } from "./Product";

@Entity("producers")
export class Producer extends BaseEntity {
  @PrimaryGeneratedColumn()
  id!: number;

  @PrimaryGeneratedColumn("uuid")
  uuid!: string;

  @Column("text")
  name!: string;

  @Column()
  canDeliver!: boolean;

  @CreateDateColumn()
  createdAt!: Date;

  @OneToMany(() => Product, (product) => product.owner)
  products: Product[];
}
