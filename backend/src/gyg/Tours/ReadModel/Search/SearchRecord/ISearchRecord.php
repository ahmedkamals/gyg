<?php

namespace gyg\Tours\ReadModel\Search\SearchRecord;

use gyg\Tours\ReadModel\Search\Strategy\ISearchStrategy;

/**
 * Interface AbstractSearchRecord
 * @package gyg\Tours\ReadModel\Search\SearchRecord
 */
interface ISearchRecord
{
    /**
     * @return ISearchStrategy
     */
    function getStrategy(): ISearchStrategy;

    /**
     * @return string
     */
    function getSearchData(): string;
}