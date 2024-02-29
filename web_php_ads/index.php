<?php

$uri = "mongodb://192.168.18.93:25001";
$client = new MongoDB\Driver\Manager($uri);

$database = "web_php_ads";
$collection = "click_log";

$document = [
    'name' =>  "John",
    'age' => 30
];

$bulk = new MongoDB\Driver\BulkWrite;
$bulk->insert($document);

$client->executeBulkWrite("$database.$collection", $bulk);

echo "Document inserted successfully!";