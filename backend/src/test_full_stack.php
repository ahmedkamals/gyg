<?php

require '../vendor/autoload.php';

use gyg\Tours\Service;
use gyg\Tours\ReadModel\Search\Factory;
use gyg\Tours\ReadModel\DataReader;
use gyg\Tours\ReadModel\Search\Strategy;
use gyg\Tours\ReadModel\Search\Criteria;
use gyg\Tours\Infrastructure\Repository;
use gyg\Tours\Exception;
use gyg\Response\JSONResponse;

try {
    $query = $_GET['query'];

    if (empty($query)) {
        throw new Exception\UnprocessableEntity('error');
    }

    $toursRepository = new Repository\ToursRepository();
    $toursSearchRecordFactory = new Factory\ToursSearchRecordFactory();
    $toursSearchDataReader = new DataReader\ToursDataReader($toursRepository);
    $toursService = new Service\ToursService($toursSearchRecordFactory, $toursSearchDataReader);

    $toursSearchRecord = $toursService->getToursSearchRecord(
        new Strategy\FindByCityStrategy(
            new Criteria\EqualToCriteria()
        ),
        $query
    );

    $data = $toursService->search($toursSearchRecord);

    (new JSONResponse($data))->render();
} catch (Exception\HttpException $e) {
    header($e->getMessage(), true, $e->getCode());
}





