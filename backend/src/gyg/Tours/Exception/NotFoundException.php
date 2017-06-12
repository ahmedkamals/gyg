<?php

namespace gyg\Tours\Exception;

/**
 * Class NotFoundException
 * @package gyg\Tours\Exception
 */
class NotFoundException extends HttpException
{
    /**
     * NotFoundException constructor.
     * 
     * @param string $message exception message.
     */
    public function __construct(string $message)
    {
        parent::__construct($message, static::NOT_FOUND);
    }
}