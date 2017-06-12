<?php


namespace gyg\Tours\Exception;

class UnprocessableEntity extends HttpException
{
    /**
     * UnprocessableEntity constructor.
     *
     * @param string $message exception message.
     */
    public function __construct(string $message)
    {
        parent::__construct($message, static::UNPROCESSABLE_ENTITY);
    }
}