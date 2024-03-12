"""Functions to keep track and alter inventory."""


def create_inventory(items: list) -> dict:
    """Create a dict that tracks the amount (count) of each element on the `items` list.

    :param items: list - list of items to create an inventory from.
    :return: dict - the inventory dictionary.
    """
    unique_items = set(items)
    inventory = {}

    for item in unique_items:
        inventory[item] = items.count(item)

    return inventory


def add_items(inventory: dict, items: list) -> dict:
    """Add or increment items in inventory using elements from the items `list`.

    :param inventory: dict - dictionary of existing inventory.
    :param items: list - list of items to update the inventory with.
    :return: dict - the inventory updated with the new items.
    """
    new_items = create_inventory(items)

    for item_key, item_value in new_items.items():
        if item_key in inventory:
            inventory[item_key] += item_value
        else:
            inventory[item_key] = item_value

    return inventory


def decrement_items(inventory: dict, items: list) -> dict:
    """Decrement items in inventory using elements from the `items` list.

    :param inventory: dict - inventory dictionary.
    :param items: list - list of items to decrement from the inventory.
    :return: dict - updated inventory with items decremented.
    """

    items_to_decrement = create_inventory(items)

    for item_key, item_value in items_to_decrement.items():
        if item_key in inventory:
            inventory[item_key] = inventory[item_key] - item_value

            if inventory[item_key] < 0:
                inventory[item_key] = 0

    return inventory


def remove_item(inventory: dict, item: str) -> dict:
    """Remove item from inventory if it matches `item` string.

    :param inventory: dict - inventory dictionary.
    :param item: str - item to remove from the inventory.
    :return: dict - updated inventory with item removed. Current inventory if item does not match.
    """

    inventory.pop(item, None)

    return inventory


def list_inventory(inventory) -> list:
    """Create a list containing all (item_name, item_count) pairs in inventory.

    :param inventory: dict - an inventory dictionary.
    :return: list of tuples - list of key, value pairs from the inventory dictionary.
    """

    list_items = []

    for key in inventory:
        if inventory[key] > 0:
            list_items.append((key, inventory[key]))

    return list_items
