<?php

namespace gyg\Tours\Service;

use gyg\Tours\ReadModel\DataReader\ToursDataReader;
use gyg\Tours\ReadModel\Search\SearchRecord\ToursSearchRecord;
use gyg\Tours\ReadModel\Search\Strategy\ISearchStrategy;
use gyg\Tours\ReadModel\Search\Factory\ToursSearchRecordFactory;

/**
 * Class ToursService
 * 
 * @package Tours\Service
 */
class ToursService
{
    /**
     * @var ToursSearchRecordFactory
     */
    private $toursSearchFactory;

    /**
     * @var ToursDataReader
     */
    private $toursDataReader;

    /**
     * ToursService constructor.
     *
     * @param ToursSearchRecordFactory $toursSearchFactory
     * @param ToursDataReader          $toursDataReader
     */
    public function __construct(ToursSearchRecordFactory $toursSearchFactory, ToursDataReader $toursDataReader)
    {
        $this->toursSearchFactory = $toursSearchFactory;
        $this->toursDataReader = $toursDataReader;
    }

    public function search(ToursSearchRecord $searchRecord): array
    {
        return $this->toursDataReader->search($searchRecord);
    }

    /**
     * @param ISearchStrategy $searchStrategy
     * @param string          $data
     *
     * @return ToursSearchRecord
     */
    public function getToursSearchRecord(ISearchStrategy $searchStrategy, string $data): ToursSearchRecord
    {
        return $this->toursSearchFactory->createToursSearchRecord($searchStrategy, $data);
    }
}