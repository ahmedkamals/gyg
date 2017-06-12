<?php

namespace gyg\Tours\ReadModel\DataReader;

use gyg\Tours\Exception\NotFoundException;
use gyg\Tours\Infrastructure\Repository\ToursRepository;
use gyg\Tours\ReadModel\Search\SearchRecord\ISearchRecord;
use gyg\Tours\ReadModel\Search\Strategy\FindByCityStrategy;

/**
 * Class ToursDataReader
 * @package gyg\Tours\ReadModel\DataReader
 */
class ToursDataReader
{
    /**
     * @var ToursRepository
     */
    private $repository;

    /**
     * ToursDataReader constructor.
     *
     * @param ToursRepository $repository
     */
    public function __construct(ToursRepository $repository)
    {
        $this->repository = $repository;
    }

    /**
     * @param ISearchRecord $searchRecord
     *
     * @return array
     * @throws NotFoundException
     */
    public function search(ISearchRecord $searchRecord): array
    {
        $data = $this->repository->get();

        $searchData = $searchRecord->getSearchData();
        $searchStrategy = $searchRecord->getStrategy();
        $searchCriteria = $searchStrategy->getCriteria();

        switch (get_class($searchStrategy))
        {
            case FindByCityStrategy::class:

                if ($searchCriteria->isEqual($searchData, 'berlin')) {
                    return $data;
                }
                break;
        }

        throw new NotFoundException('error');
    }
}