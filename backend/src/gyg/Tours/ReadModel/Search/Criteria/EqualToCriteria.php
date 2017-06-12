<?php

namespace gyg\Tours\ReadModel\Search\Criteria;

/**
 * Class EqualToCriteria
 * @package gyg\Tours\ReadModel\Search\Criteria
 */
class EqualToCriteria implements ISearchCriteria
{
    /**
     * @inheritdoc
     */
    public function isEqual($first, $second, bool $strict = false): bool
    {
        if ($strict) {
            return $first === $second;
        }
        
        return $first == $second;
    }
}