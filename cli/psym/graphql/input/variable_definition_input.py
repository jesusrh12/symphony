#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from gql_client.runtime.enum_utils import enum_field_metadata
from ..enum.variable_type import VariableType


@dataclass(frozen=True)
class VariableDefinitionInput(DataClassJsonMixin):
    key: str
    type: VariableType = _field(metadata=enum_field_metadata(VariableType))
    choices: List[str]
    mandatory: Optional[bool] = None
    multipleValues: Optional[bool] = None
    defaultValue: Optional[str] = None
