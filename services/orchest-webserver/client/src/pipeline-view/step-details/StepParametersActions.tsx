import { IconButton } from "@/components/common/IconButton";
import MoreHorizOutlinedIcon from "@mui/icons-material/MoreHorizOutlined";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import MenuList from "@mui/material/MenuList";
import Stack from "@mui/material/Stack";
import ToggleButton from "@mui/material/ToggleButton";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import React from "react";
import { useStepDetailsContext } from "./StepDetailsContext";
import { useOpenSchemaFile } from "./useOpenSchemaFile";

export type StepParameterViewingMode = "json" | "form";

type StepParametersActionsProps = {
  viewingMode: StepParameterViewingMode;
  setViewingMode: (
    value:
      | StepParameterViewingMode
      | ((value: StepParameterViewingMode) => StepParameterViewingMode)
  ) => void;
};

export const StepParametersActions = ({
  viewingMode,
  setViewingMode,
}: StepParametersActionsProps) => {
  const moreOptionsButtonRef = React.useRef<HTMLButtonElement | null>(null);

  const [isMoreOptionsOpen, setIsMoreOptionsOpen] = React.useState(false);
  const openMoreOptions = () => setIsMoreOptionsOpen(true);
  const closeMoreOptions = () => setIsMoreOptionsOpen(false);

  const { parameterSchema, parameterUiSchema } = useStepDetailsContext();
  const { openSchemaFile } = useOpenSchemaFile();

  return (
    <Stack direction="row" justifyContent="space-between">
      <ToggleButtonGroup
        exclusive
        size="small"
        aria-label="Step parameters viewing mode"
        value={viewingMode}
        onChange={(e, value) => setViewingMode(value)}
      >
        <ToggleButton value="json" disabled={viewingMode === "json"}>
          Json
        </ToggleButton>
        <ToggleButton value="form" disabled={viewingMode === "form"}>
          Form
        </ToggleButton>
      </ToggleButtonGroup>
      <IconButton
        title="More options"
        ref={moreOptionsButtonRef}
        onClick={openMoreOptions}
        id="more-options-button"
      >
        <MoreHorizOutlinedIcon />
      </IconButton>
      <Menu
        id="running-services-menu"
        anchorEl={moreOptionsButtonRef.current}
        open={isMoreOptionsOpen}
        onClose={closeMoreOptions}
        MenuListProps={{
          dense: true,
          "aria-labelledby": "more-options-button",
        }}
      >
        <MenuList dense>
          <MenuItem
            onClick={(e) => openSchemaFile(e, ".schema.json")}
            onAuxClick={(e) => openSchemaFile(e, ".schema.json")}
          >
            {`${parameterSchema ? "Edit" : "New"} schema file`}
          </MenuItem>
          <MenuItem
            disabled={!parameterSchema}
            onClick={(e) => openSchemaFile(e, ".uischema.json")}
            onAuxClick={(e) => openSchemaFile(e, ".uischema.json")}
          >
            {`${parameterUiSchema ? "Edit" : "New"} UI schema file`}
          </MenuItem>
        </MenuList>
      </Menu>
    </Stack>
  );
};
