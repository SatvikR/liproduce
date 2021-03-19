import dotenv from "dotenv";
import path from "path";
import { ConnectionOptions } from "typeorm";
import { Producer } from "./entities/Producer";
import { Product } from "./entities/Product";

dotenv.config({
  path: process.env.NODE_ENV === "production" ? ".env.prod" : ".env.dev",
});

export const CORS_ORIGIN: string | undefined = process.env.CORS_ORIGIN;
export const PORT = process.env.PORT;
export const DB_CONFIG: ConnectionOptions = {
  type: "postgres",
  database: process.env.DB,
  username: process.env.DB_USERNAME,
  password: process.env.DB_PASSWORD,
  logging: process.env.NODE_ENV !== "production",
  synchronize: process.env.NODE_ENV !== "production",
  entities: [Product, Producer],
  migrations: [path.join(__dirname, "./migrations/*")],
};
