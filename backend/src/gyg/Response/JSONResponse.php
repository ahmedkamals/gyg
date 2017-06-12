<?php

namespace gyg\Response;

/**
 * Class JSONResponse
 * @package gyg\Response
 */
class JSONResponse implements IResponse
{
    /**
     * @var array
     */
    private $content;

    /**
     * JSONResponse constructor.
     *
     * @param array $content
     */
    public function __construct(array $content)
    {
        $this->content = $content;
    }

    /**
     * @return string
     */
    public function getContent(): string
    {
        return json_encode($this->content);
    }

    public function render()
    {
        header('content-type: application/json');
        echo $this->getContent();
    }
}