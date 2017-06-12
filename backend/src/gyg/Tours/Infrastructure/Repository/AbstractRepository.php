<?php

namespace gyg\Tours\Infrastructure\Repository;

/**
 * Class AbstractRepository
 * @package gyg\Tours\Infrastructure\Repository
 */
Abstract class AbstractRepository
{
    /**
     * Returns array of data.
     * 
     * @return array
     */
    public abstract function get(): array;
}