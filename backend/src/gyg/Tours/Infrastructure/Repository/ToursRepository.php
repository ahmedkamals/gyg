<?php

namespace gyg\Tours\Infrastructure\Repository;

/**
 * Class ToursRepository
 * @package gyg\Tours\Infrastructure\Repository
 */
class ToursRepository extends AbstractRepository
{
    /**
     * @inheritdoc
     */
    public function get(): array
    {
        return [
            'tours' => [
                [
                    'id' => 1,
                    'title' => 'German Tour: Parliament Quarter \u0026amp; Reichstag glass dome',
                    'price' => (float) 14,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 2,
                    'title' => 'Skip the Line: Pergamon Museum Berlin Tickets',
                    'price' => (float) 21,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 3,
                    'title' => 'Berlin WelcomeCard: Transport, Discounts \u0026amp; Guide Book',
                    'price' => (float) 10,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 4,
                    'title' => 'Skip the Line: Berlin TV Tower Ticket',
                    'price' => (float) 143,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 5,
                    'title' => '1-Hour City Tour by Boat with Guaranteed Seating',
                    'price' => (float) 14,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 6,
                    'title' => 'Berlin Hop-on Hop-off Bus Tour: 1- or 2-Day Ticket',
                    'price' => (float) 210,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 7,
                    'title' => 'German Tour: Reichstag with Plenary Chamber \u0026amp; Cuppola',
                    'price' => (float) 59,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 8,
                    'title' => 'Berlin: 2.5-Hour Boat Tour Along the River Spree',
                    'price' => (float) 41,
                    'currency' => '$',
                    'isSpecialOffer' => true
                ],
                [
                    'id' => 9,
                    'title' => 'Museum Pass Berlin: 3-Day Entry to Over 40 Top Museums',
                    'price' => (float) 14,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 10,
                    'title' => 'Reichstag: Rooftop Coffee Break at K\u00e4fer',
                    'price' => (float) 50,
                    'currency' => '$',
                    'isSpecialOffer' => true
                ],
                [
                    'id' => 11,
                    'title' => 'Skip the Line: Neues Museum Berlin Tickets',
                    'price' => (float) 46,
                    'currency' => '$',
                    'isSpecialOffer' => true
                ],
                [
                    'id' => 12,
                    'title' => 'German Tour: Reichstag with dome Chamber \u0026amp;',
                    'price' => (float) 87,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 13,
                    'title' => 'Berlin Hop-on Hop-off Bus Tour with Live Commentary',
                    'price' => (float) 8,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ],
                [
                    'id' => 14,
                    'title' => 'Skip the Line: TV Tower Early Bird Tickets',
                    'price' => (float) 140,
                    'currency' => '$',
                    'isSpecialOffer' => false
                ]
            ]
        ];
    }
}