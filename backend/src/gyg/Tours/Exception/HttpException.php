<?php

namespace gyg\Tours\Exception;

/**
 * Class HttpException
 * @package gyg\Tours\Exception
 */
class HttpException extends \Exception
{
    const NOT_FOUND = 404;
    const UNPROCESSABLE_ENTITY = 422;
}