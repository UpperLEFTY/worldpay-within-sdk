"""Abstract class for event listeners to inherit and implement.

For examples, see
https://github.com/WPTechInnovation/worldpay-within-sdk/tree/feature/python-wrapper/wrappers/python/wpwithin_python/examples
"""


class AbstractEventListener(object):
    """Abstract class with methods required by WPWithin Callback Service.

    For examples, see
    https://github.com/WPTechInnovation/worldpay-within-sdk/tree/feature/python-wrapper/wrappers/python/wpwithin_python/examples
    """

    def beginServiceDelivery(self, service_id, service_delivery_token, units_to_supply):
        """Run when service delivery starts.

        Must be called by consumer.
        """
        raise NotImplementedError("You should implement this method in a derived class.")

    def endServiceDelivery(self, service_id, service_delivery_token, units_received):
        """Run when service delivery starts.

        Must be called by consumer.
        """
        raise NotImplementedError("You should implement this method in a derived class.")
