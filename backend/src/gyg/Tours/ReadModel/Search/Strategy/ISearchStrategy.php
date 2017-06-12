<?php

namespace gyg\Tours\ReadModel\Search\Strategy;

use gyg\Tours\ReadModel\Search\Criteria\ISearchCriteria;

/**
 * Interface ISearchStrategy
 * @package gyg\Tours\ReadModel\Search\Strategy
 */
interface ISearchStrategy
{
    public function getCriteria(): ISearchCriteria;
}