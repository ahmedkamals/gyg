<?php

namespace gyg\Tours\ReadModel\Search\SearchRecord;

use gyg\Tours\ReadModel\Search\Strategy\ISearchStrategy;

/**
 * Class ToursSearchRecord
 * @package gyg\Tours\ReadModel\Search\SearchRecord
 */
class ToursSearchRecord implements ISearchRecord
{
    /**
     * @var ISearchStrategy
     */
    private $strategy;

    /**
     * @var string
     */
    private $searchData;

    /**
     * ToursSearchRecord constructor.
     *
     * @param ISearchStrategy $strategy
     * @param string           $data
     */
    public function __construct(ISearchStrategy $strategy, string $data)
    {
        $this->strategy = $strategy;
        $this->searchData = $data;
    }

    /**
     * @return string
     */
    public function getSearchData(): string
    {
        return $this->searchData;
    }

    /**
     * @return ISearchStrategy
     */
    public function getStrategy(): ISearchStrategy
    {
        return $this->strategy;
    }
}