<?php

namespace gyg\Tours\ReadModel\Search\Strategy;

use gyg\Tours\ReadModel\Search\Criteria\ISearchCriteria;

/**
 * Class FindByCityStrategy
 * @package gyg\Tours\ReadModel\Search\Strategy
 */
class FindByCityStrategy implements ISearchStrategy
{
    /**
     * @var ISearchCriteria
     */
    private $criteria;

    /**
     * FindByCityStrategy constructor.
     *
     * @param ISearchCriteria $criteria
     */
    public function __construct(ISearchCriteria $criteria)
    {
        $this->criteria = $criteria;
    }

    /**
     * @inheritdoc
     */
    public function getCriteria(): ISearchCriteria
    {
        return $this->criteria;
    }
}