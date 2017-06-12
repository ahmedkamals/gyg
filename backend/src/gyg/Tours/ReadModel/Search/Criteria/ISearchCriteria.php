<?php

namespace gyg\Tours\ReadModel\Search\Criteria;

/**
 * Interface ISearchCriteria
 * @package gyg\Tours\ReadModel\Search\Criteria
 */
interface ISearchCriteria
{
    /**
     * Checks if two values are equal.
     *
     * @param mixed $first
     * @param mixed $second
     * @param bool  $strict
     *
     * @return bool
     */
    public function isEqual($first, $second, bool $strict = false): bool;
}