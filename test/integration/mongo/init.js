db.createUser({
    user: 'application_user',
    pwd: 'application_pass',
    roles: [
        {
            role: 'dbOwner',
            db: 'beverageDeliveryManagerDB',
        },
    ],
});

db.pdvs.createIndex({"address": "2dsphere"});
db.pdvs.createIndex({"coverageArea": "2dsphere"});