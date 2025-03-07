
import React from 'react'
// import { CSSTransition } from 'react-transition-group'
import {
  // Classes,
  Icon,
  Spinner,
  Colors,
  Tooltip,
  Position,
  Intent,
  // Alignment
} from '@blueprintjs/core'
import dayjs from '@/utils/time'

const TaskActivity = (props) => {
  const { activePipeline } = props

  return (
    <>

      <div
        className='pipeline-task-activity' style={{
          padding: '20px',
          overflow: 'hidden',
          textOverflow: 'ellipsis'
        }}
      >
        {activePipeline?.ID && activePipeline.tasks && activePipeline.tasks.map((t, tIdx) => (
          <div
            className='pipeline-task-row'
            key={`pipeline-task-key-${tIdx}`}
            style={{
              display: 'flex',
              padding: '4px 0',
              justifyContent: 'space-between',
              fontSize: '12px',
              opacity: t.status === 'TASK_CREATED' ? 0.7 : 1
            }}
          >
            <div style={{ display: 'flex', justifyContent: 'center', paddingRight: '8px', width: '32px', minWidth: '32px' }}>
              {t.status === 'TASK_COMPLETED' && (
                <Tooltip content={`Task Complete [STAGE ${t.pipelineRow}]`} position={Position.TOP} intent={Intent.SUCCESS}>
                  <Icon icon='small-tick' size={18} color={Colors.GREEN5} style={{ marginLeft: '0' }} />
                </Tooltip>
              )}
              {t.status === 'TASK_FAILED' && (
                <Tooltip content={`Task Failed [STAGE ${t.pipelineRow}]`} position={Position.TOP} intent={Intent.PRIMARY}>
                  <Icon icon='warning-sign' size={14} color={Colors.RED5} style={{ marginLeft: '0', marginBottom: '3px' }} />
                </Tooltip>
              )}
              {t.status === 'TASK_RUNNING' && (
                <Tooltip content={`Task Running [STAGE ${t.pipelineRow}]`} position={Position.TOP}>
                  <Spinner
                    className='task-spinner'
                    size={14}
                    intent={t.status === 'TASK_COMPLETED' ? 'success' : 'warning'}
                    value={t.status === 'TASK_COMPLETED' ? 1 : t.progress}
                  />
                </Tooltip>
              )}
              {t.status === 'TASK_CREATED' && (
                <Tooltip content={`Task Created (Pending) [STAGE ${t.pipelineRow}]`} position={Position.TOP}>
                  <Icon icon='pause' size={14} color={Colors.GRAY3} style={{ marginLeft: '0', marginBottom: '3px' }} />
                </Tooltip>
              )}
            </div>
            <div
              className='pipeline-task-cell-name'
              style={{ padding: '0 8px', minWidth: '100px', display: 'flex', justifyContent: 'space-between' }}
            >
              <strong
                className='task-plugin-name'
                style={{ overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}
              >
                {t.plugin}
              </strong>

            </div>
            <div
              className='pipeline-task-cell-settings'
              style={{
                padding: '0 8px',
                display: 'flex',
                width: '20%',
                minWidth: '20%',
                // whiteSpace: 'nowrap',
                justifyContent: 'flex-start',
                // overflow: 'hidden',
                // textOverflow: 'ellipsis',
                flexGrow: 1
              }}
            >
              {t.plugin !== 'jenkins' && (
                <div style={{ overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>
                  <span style={{ color: Colors.GRAY2 }}>
                    <Icon icon='link' size={8} style={{ marginBottom: '3px' }} /> {t.options[Object.keys(t.options)[0]]}
                  </span>
                  {t.plugin === 'github' && (
                    <span style={{ fontWeight: 60 }}>/{t.options[Object.keys(t.options)[1]]}</span>
                  )}
                </div>
              )}
            </div>
            <div
              className='pipeline-task-cell-duration'
              style={{
                padding: '0',
                minWidth: '80px',
                // whiteSpace: 'nowrap',
                textAlign: 'right'
              }}
            >
              <span style={{ whiteSpace: 'nowrap' }}>
                {(() => {
                  let statusRelativeTime = dayjs(t.CreatedAt).toNow(true)
                  switch (t.status) {
                    case 'TASK_COMPLETED':
                    case 'TASK_FAILED':
                      statusRelativeTime = dayjs(t.UpdatedAt).from(t.CreatedAt, true)
                      break
                    case 'TASK_RUNNING':
                    default:
                      statusRelativeTime = dayjs(t.CreatedAt).toNow(true)
                      break
                  }
                  return statusRelativeTime
                })()}
              </span>
            </div>
            <div
              className='pipeline-task-cell-progress'
              style={{
                padding: '0 8px',
                minWidth: '100px',
                textAlign: 'right'
              }}
            >
              <span style={{ fontWeight: t.status === 'TASK_COMPLETED' ? 800 : 600 }}>
                {Number(t.status === 'TASK_COMPLETED' ? 100 : (t.progress / 1) * 100).toFixed(2)}%
              </span>
            </div>
            <div
              className='pipeline-task-cell-message'
              style={{ display: 'flex', flexGrow: 1, width: '64%' }}
            >
              {t.message && (
                <div style={{ width: '98%', whiteSpace: 'nowrap', overflow: 'hidden', textOverflow: 'ellipsis' }}>
                  <span style={{ color: t.status === 'TASK_FAILED' ? Colors.RED4 : Colors.GRAY3, paddingLeft: '10px' }}>
                    {t.message}
                  </span>
                </div>
              )}
            </div>
          </div>
        ))}
      </div>
    </>
  )
}

export default TaskActivity
