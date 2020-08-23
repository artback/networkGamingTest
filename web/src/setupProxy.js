const proxy = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    proxy.createProxyMiddleware("/backend", {
      target: "http://localhost:8080",
        ws: true
    })
  )
};
