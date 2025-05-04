// insertData.js

const { MongoClient, ObjectId } = require('mongodb');
const { faker } = require('@faker-js/faker');

// URL de conexión (modifica con tu usuario, password, host y base de datos si es necesario)
const uri = 'mongodb://localhost:27017';
const dbName = 'tododatabase';
const collectionName = 'tasks';

async function insertFakeData() {
    const client = new MongoClient(uri);

    try {
        await client.connect();
        const db = client.db(dbName);
        const collection = db.collection(collectionName);

        const documents = [];

        for (let i = 0; i < 10; i++) {
            const _id = new ObjectId();
            documents.push({
                _id: _id,
                id: _id.toString(),
                title: faker.word.words(2),
                description: faker.lorem.sentence(),
                completed: faker.datatype.boolean(),
                created_at: new Date().toISOString()
            });
        }

        const result = await collection.insertMany(documents);
        console.log(`${result.insertedCount} documentos insertados exitosamente.`);
    } catch (error) {
        console.error('Error al insertar los documentos:', error);
    } finally {
        await client.close();
    }
}

insertFakeData();
