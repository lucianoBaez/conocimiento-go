db.auth('admin-user', 'admin-password')
db = db.getSiblingDB('drivers_mongo')
db.createUser(
    {
        user: "dunauser",
        pwd: "password",
        roles: [
            {
                "role": "readWrite",
                "db": "drivers_mongo"
            }
        ]
    }
);

db.auth('dunauser', 'password')

// Adding 100 drivers
for(let driverIndex = 0; driverIndex < 100; driverIndex++) {
    db.drivers.insert( {
        "name": "Name" + "-" + driverIndex,
        "lastname": "Lastname" + "-" + driverIndex,
        "email": "lutiano2@gmail.com",
        "location": Math.floor(Math.random() * 5),
        "password": "123456",
        "createdAt": "2021-04-09T12:08:54.674Z",
        "updatedAt": "2021-04-09T12:08:54.674Z"
    });
}