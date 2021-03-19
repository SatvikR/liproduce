import { ApolloServer } from "apollo-server-express";
import cors from "cors";
import express from "express";
import "reflect-metadata";
import { buildSchema } from "type-graphql";
import { createConnection } from "typeorm";
import { CORS_ORIGIN, DB_CONFIG } from "./constants";
import { HelloResolver } from "./resolvers/HelloWorld";

(async () => {
  await createConnection(DB_CONFIG);

  const app = express();

  app.use(
    cors({
      origin: CORS_ORIGIN,
    })
  );

  const apolloServer = new ApolloServer({
    schema: await buildSchema({
      resolvers: [HelloResolver],
      validate: false,
    }),
  });

  apolloServer.applyMiddleware({
    app,
    cors: false,
  });

  app.listen(4000, () => {
    console.log("server started");
  });
})().catch((err) => console.log(err));
