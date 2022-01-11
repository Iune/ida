from ida.contest import Country


def test_contained_in_regex():
    country = Country("unitedkingdom", ["UK", "United Kingdom"])

    lines = [
        "12 uk",
        "12 ukraine",
        "12 UK",
        "12 Ukraine",
        "12 United Kingdom",
        "12 Srbuk",
    ]

    present = [country.contained_in_regex(line) for line in lines]
    assert present == [True, False, True, False, True, False]
