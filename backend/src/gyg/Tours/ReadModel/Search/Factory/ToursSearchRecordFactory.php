<?php

namespace gyg\Tours\ReadModel\Search\Factory;

use gyg\Tours\ReadModel\Search\SearchRecord\ToursSearchRecord;
use gyg\Tours\ReadModel\Search\Strategy\ISearchStrategy;

/**
 * Class ToursSearchRecordFactory
 * @package gyg\Tours\ReadModel\Search\Factory
 */
class ToursSearchRecordFactory
{
    /**
     * Creates ToursSearchRecord
     * 
     * @param ISearchStrategy $strategy
     * @param string          $searchData
     * 
     * @return ToursSearchRecord
     */
    public function createToursSearchRecord(ISearchStrategy $strategy, string $searchData): ToursSearchRecord
    {
        return new ToursSearchRecord($strategy, $searchData);
    }
}