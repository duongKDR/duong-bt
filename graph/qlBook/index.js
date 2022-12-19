const express = require("express");
const { ApolloServer } = require("apollo-server-express");
const http = require("http");

const app = express();

const typeDefs = require('../qlBook/shema/shema');
const resolvers = require('../qlBook/Resolvers/resolvers');

let apolloServer = null;
async function startServer() {
  apolloServer = new ApolloServer({
    typeDefs,
    resolvers,
  });
  await apolloServer.start();
  apolloServer.applyMiddleware({ app });
}
startServer();
const httpserver = http.createServer(app);

app.get("/rest", function (req, res) {
  res.json({ data: "api working" });
});

app.listen(4000, function () {
  console.log(`server running on port 4000`);
  console.log(`Server ready at http://localhost:4000${apolloServer.graphqlPath}`);
});
