#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from ...config import custom_scalars, datetime
from gql_client.runtime.variables import encode_variables
from gql import gql, Client
from gql.transport.exceptions import TransportQueryError
from functools import partial
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin, config

from ..fragment.equipment_port_type import EquipmentPortTypeFragment, QUERY as EquipmentPortTypeFragmentQuery

from ..input.add_equipment_port_type_input import AddEquipmentPortTypeInput


# fmt: off
QUERY: List[str] = EquipmentPortTypeFragmentQuery + ["""
mutation AddEquipmentPortTypeMutation($input: AddEquipmentPortTypeInput!) {
  addEquipmentPortType(input: $input) {
    ...EquipmentPortTypeFragment
  }
}

"""
]


class AddEquipmentPortTypeMutation:
    @dataclass(frozen=True)
    class AddEquipmentPortTypeMutationData(DataClassJsonMixin):
        @dataclass(frozen=True)
        class EquipmentPortType(EquipmentPortTypeFragment):
            pass

        addEquipmentPortType: EquipmentPortType

    # fmt: off
    @classmethod
    def execute(cls, client: Client, input: AddEquipmentPortTypeInput) -> AddEquipmentPortTypeMutationData.EquipmentPortType:
        variables: Dict[str, Any] = {"input": input}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = client.execute(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.AddEquipmentPortTypeMutationData.from_dict(response_text)
        return res.addEquipmentPortType

    # fmt: off
    @classmethod
    async def execute_async(cls, client: Client, input: AddEquipmentPortTypeInput) -> AddEquipmentPortTypeMutationData.EquipmentPortType:
        variables: Dict[str, Any] = {"input": input}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = await client.execute_async(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.AddEquipmentPortTypeMutationData.from_dict(response_text)
        return res.addEquipmentPortType
