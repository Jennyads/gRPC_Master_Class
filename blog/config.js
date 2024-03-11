module.exports = {
    options: {
        useBasicAuth: true,
        useNewUrlParser: true,
        useUnifiedTopology: true,
        allowMethodOverride: false,
        allowHttpMethod: false,
        allowInsecureHTTP: false,
        listDatabases: false, // Desabilitar o comando listDatabases
        sslCert: '/etc/ssl/mongodb.pem',
    },
};
